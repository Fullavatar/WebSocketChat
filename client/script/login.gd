extends Control

func serverConnect(address: String, port: String) -> void:
	if address and port:
		var ws = WebSocketPeer.new()
		if ws.connect_to_url("ws://%s:%s/ws" % [address, port]) == OK:
			
			while ws.get_ready_state() != ws.STATE_OPEN:
				ws.poll()
				await get_tree().process_frame

			var ping = load("res://script/message_template/message_resource/ping.tres")
			ws.send_text(ping.to_json())

			var timeout = 10.0
			var start_time = Time.get_ticks_msec() / 1000.0

			while ws.get_ready_state() == WebSocketPeer.STATE_OPEN:
				ws.poll()
				if ws.get_available_packet_count() > 0:
					var packet = ws.get_packet()
					var json_string = packet.get_string_from_utf8()
					var response_data = JSON.parse_string(json_string)
					if response_data["Type"] != "Ping":
						#TODO: add logs
						ws.close()
						return
					authenticate(address, port)
					return
				if (Time.get_ticks_msec() / 1000.0) - start_time >= timeout:
					ws.close()
					return
				await get_tree().process_frame
			print("Connection closed")


func authenticate(address, port):
	var chat = preload("res://scene/main.tscn").instantiate()
	chat.address = "ws://%s:%s/ws"% [address, port]
	get_tree().get_root().add_child(chat)
	hide()
	queue_free()
