extends Control

func _on_connect_pressed() -> void:
	updateConnectionUIState(false)
	get_tree().call_group(
		"login",
		"serverConnectionHandler",
		$Address.text,
		$Port.get_line_edit().text,
	)


func updateConnectionUIState(can_edit:bool):
	$Address.editable = can_edit
	$Port.editable = can_edit
	$Connect.disabled = !can_edit
