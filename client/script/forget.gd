extends Control


func _on_back_pressed() -> void:
	get_tree().call_group(
		"login",
		"swapUIAuthenticate",
		Global.AuthenticateUIType.AUTHENTICATE,
	)
