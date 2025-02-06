extends Control


func _on_connect_pressed() -> void:
	var address = $Address.text
	var port = $Port.text
	
	if address and port:
		var chat = preload("res://scene/main.tscn").instantiate()
		chat.address = "ws://%s:%s/ws"% [address, port]
		get_tree().get_root().add_child(chat)
		hide()
		queue_free()
