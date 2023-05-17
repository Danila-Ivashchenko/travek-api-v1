<?php
	require_once 'login.php';
	$conn = new mysqli($hn, $un, $pw, $db);
	if (isset($_POST['name'])){
		$name = $_POST['name'];

		$new_url = 'images/';
		$query = "SELECT url FROM images WHERE name = $name";
		$result = $conn->query($query);
		$new_url .= $result->fetch_array(MYSQLI_ASSOC)['url'];
		echo $new_url;
	}
?>