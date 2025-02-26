extends Control

var ws:WebSocketPeer
var address:String
var port:String
var timeout := 10.0
@onready var logger:Logger = $Log

func serverConnectionHandler(addr:String, prt:String):
	address = addr
	port = prt
	if !addressValidation():
		return

	handle_connection_result(Global.LogType.CONNECTING)
	
	var connection_validation = await serverConnectionValidation()
	handle_connection_result(connection_validation)
	if connection_validation != Global.LogType.CONNECTION_ESTABLISHED:
		return
	
	var response_validation = await serverResponseValidation()
	handle_connection_result(response_validation)


func swapUIAuthenticate(swap_to:Global.AuthenticateUIType):
	match swap_to:
		Global.AuthenticateUIType.NONE:
			$Authenticate.hide()
			$SignUp.hide()
			$Forget.hide()
		Global.AuthenticateUIType.AUTHENTICATE: 
			$Authenticate.show()
			$SignUp.hide()
			$Forget.hide()
		Global.AuthenticateUIType.SIGNUP: 
			$Authenticate.hide()
			$SignUp.show()
			$Forget.hide()
		Global.AuthenticateUIType.FORGET: 
			$Authenticate.hide()
			$SignUp.hide()
			$Forget.show()


func handle_connection_result(log_type: Global.LogType):
	logger.write(log_type)
	match log_type:
		Global.LogType.NOT_VALID_ADDRESS:
			$Server.updateConnectionUIState(true)
		Global.LogType.CONNECTION_FAILED:
			ws.close()
			$Server.updateConnectionUIState(true)
		Global.LogType.CONNECTION_TIMEOUT:
			ws.close()
			$Server.updateConnectionUIState(true)
		Global.LogType.CANNOT_PARSE_RESPONSE:
			ws.close()
			$Server.updateConnectionUIState(true)
		Global.LogType.CONNECTION_SUCCESS:
			swapUIAuthenticate(Global.AuthenticateUIType.AUTHENTICATE)


func addressValidation()-> bool:
	if !address:
		handle_connection_result(Global.LogType.NOT_VALID_ADDRESS)
		return false
	return true


func serverConnectionValidation()-> Global.LogType:
	ws = WebSocketPeer.new()
	if ws.connect_to_url("ws://" + address + ":" + port + "/ws") != OK:
		return Global.LogType.CONNECTION_FAILED
		
	var connection_time := Time.get_ticks_msec() / 1000.0
	while ws.get_ready_state() != WebSocketPeer.STATE_OPEN:
		ws.poll()
		if ws.get_ready_state() == WebSocketPeer.STATE_CLOSED:
			return Global.LogType.CONNECTION_FAILED
			
		if ws.get_ready_state() == WebSocketPeer.STATE_CONNECTING:
			if Time.get_ticks_msec() / 1000.0 - connection_time > timeout:
				return Global.LogType.CONNECTION_TIMEOUT
			
		await get_tree().process_frame
		
	return Global.LogType.CONNECTION_ESTABLISHED


func serverResponseValidation()-> Global.LogType:
	var ping := load("res://script/message_template/message_resource/ping.tres")
	var response_time := Time.get_ticks_msec() / 1000.0
	ws.send_text(ping.to_json())
	
	while ws.get_ready_state() == WebSocketPeer.STATE_OPEN:
		ws.poll()
		if Time.get_ticks_msec() / 1000.0 - response_time > timeout:
			return Global.LogType.CONNECTION_TIMEOUT
		
		if ws.get_available_packet_count() > 0:
			var packet := ws.get_packet()
			var message := JSON.new()
			var parsed := message.parse(packet.get_string_from_utf8())
			if parsed != OK:
				return Global.LogType.CANNOT_PARSE_RESPONSE
			
			if message.data["Type"] == "Ping":
				return Global.LogType.CONNECTION_SUCCESS
		
		await get_tree().process_frame
	return Global.LogType.CONNECTION_FAILED
