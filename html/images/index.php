<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
</head>
<body>
<?php
	require_once 'login.php';
	$conn = new mysqli($hn, $un, $pw, $db);
	if (!$conn) die('Connection error');
	$query = "SELECT * FROM images WHERE url = 'fruits'";
	$result = $conn->query($query);
?>
	<form action="addressing.php">
	<?php
		for ($i = 0; $i < $result->num_rows; $i++){
			$row = $result->fetch_array(MYSQLI_ASSOC);
			$name = $row['name'];
			$url = $row['url'];
			echo <<<_END
				<label for='$i'><img src='img/$name' alter='$name'>
				<inpyt type='radiobutton' id='$i'>
				</label>
			_END;
		}
	?>
	
	</form>
</body>
</html>

