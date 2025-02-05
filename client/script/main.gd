extends Control

var address:String
var ws:WebSocketPeer
var username_set = false


func _ready() -> void:
	$Username.cancel_ime()
	$Message.cancel_ime()
	
	ws = WebSocketPeer.new()
	var err = FAILED
	while err != OK:
		err = ws.connect_to_url(address)


func _process(_delta: float) -> void:
	if ws == null:
		return
	
	ws.poll()
	
	match ws.get_ready_state():
		ws.STATE_OPEN:
			processMessages()
		ws.STATE_CLOSED:
			print("Connexion lost, trying reconnexion...")
			reconnexion()


func processMessages():
	while ws.get_available_packet_count() > 0:
		var raw_message = ws.get_packet().get_string_from_utf8()
		var parsed_message = parseReceivedMessage(raw_message)
		addMessage(parsed_message.keys()[0], parsed_message.values()[0])


func reconnexion():
	await get_tree().create_timer(2.0).timeout
	ws.connect_to_url(address)


func randomName() -> String:
	return "User" + str(randi() % 10000)


func parseReceivedMessage(received_message:String)-> Dictionary[String, String]:
	var parts = received_message.split(": ", true, 1)
	var message:Dictionary[String, String]
	message[parts[0]] = parts[1]
	return message


func addMessage(sender:String, body:String):
	var final_message
	if sender == $Username.text:
		final_message = (
			"[right]" +
				"[color=GREEN_YELLOW]" +
					sender.capitalize() + " : " +
				"[/color]\n" +
				"[color=POWDER_BLUE]" +
					"\t" + body +
				"[/color]" +
			"[/right]\n")
	else:
		final_message = (
			"[left]" +
				"[color=TOMATO]" +
					sender.capitalize() + " : " +
				"[/color] \n" +
				"[color=WHITE_SMOKE]" +
					"\t" + body +
				"[/color]" +
			"[/left]\n")
	
	$BG/Chat.append_text(final_message)


func _on_username_text_submitted(new_text: String) -> void:
	while ws.get_ready_state() != ws.STATE_OPEN:
		await get_tree().create_timer(0.1).timeout
	
	if new_text == "":
		new_text = randomName()
		$Username.text = new_text
	var err = ws.send_text(new_text)
	username_set = true
	if err == OK:
		$Username.release_focus()
		$Message.editable = true
		$Username.editable = false


func _on_message_text_submitted(new_text: String) -> void:
	if new_text != "":
		var err = ws.send_text(new_text)
		if err == OK:
			$Message.clear()
	$Message.release_focus()
	$Message.grab_focus.call_deferred()



func _on_button_pressed() -> void:
	if ws and ws.get_ready_state() == ws.STATE_OPEN:
		if username_set == false:
			ws.send_text("")
		ws.close(1000, "Client quitting")
		while ws.get_ready_state() != ws.STATE_CLOSED:
			ws.poll()
	get_tree().quit()
