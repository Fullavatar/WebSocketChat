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
					if packet.size() > 0:
						authenticate(address, port)
						return
				if (Time.get_ticks_msec() / 1000.0) - start_time >= timeout:
					print("Timeout atteint")
					return
				await get_tree().process_frame  # Laisser Godot avancer les frames


func authenticate(address, port):
	var chat = preload("res://scene/main.tscn").instantiate()
	chat.address = "ws://%s:%s/ws"% [address, port]
	get_tree().get_root().add_child(chat)
	hide()
	queue_free()
