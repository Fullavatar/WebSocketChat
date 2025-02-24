extends Control

@onready var logger:Logger = $Log

func serverConnect(address:String, port:String):
	if address == "" or port == "":
		logger.write(Global.LogType.VALIDADDRESS)
		return $Server.updateConnectionUIState(true)

	logger.write(Global.LogType.CONNECTING)
	var ws := WebSocketPeer.new()
	if ws.connect_to_url("ws://" + address + ":" + port + "/ws") != OK:
		logger.write(Global.LogType.FAILEDCONNECTION)
		return $Server.updateConnectionUIState(true)
	
	while ws.get_ready_state() != WebSocketPeer.STATE_OPEN:
		ws.poll()
		if ws.get_ready_state() == WebSocketPeer.STATE_CLOSED:
			logger.write(Global.LogType.FAILEDCONNECTION)
			ws.close()
			return $Server.updateConnectionUIState(true)
		await get_tree().process_frame
	
	var ping := load("res://script/message_template/message_resource/ping.tres")
	var timeout := 10.0
	var start_time := Time.get_ticks_msec() / 1000.0
	ws.send_text(ping.to_json())
	while ws.get_ready_state() == WebSocketPeer.STATE_OPEN:
		ws.poll()
		if Time.get_ticks_msec() / 1000.0 - start_time > timeout:
			logger.write(Global.LogType.TIMEOUT)
			ws.close()
			return $Server.updateConnectionUIState(true)
		
		if ws.get_available_packet_count() > 0:
			var packet := ws.get_packet()
			var message := JSON.new()
			var parsed := message.parse(packet.get_string_from_utf8())
			if parsed != OK:
				logger.write(Global.LogType.CANNOTPARSE)
				ws.close()
				return $Server.updateConnectionUIState(true)
			
			if message.data["Type"] == "Ping":
				logger.write(Global.LogType.CONNECTIONSUCCESS)
				return enableAuthentication()
		
		await get_tree().process_frame


func enableAuthentication():
	$Authenticate.visible = true
