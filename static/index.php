<!DOCTYPE html>
<html lang="en-US">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width" />
    <title>Random Episode</title>
    <link rel="stylesheet" type="text/css" href="index.css" /> 
  </head>
  <body>
    <?php
        $services = array_diff(scandir('/var/episodes/'), array('..', '.'));
        foreach($services as $service) {
            echo "<div class=\"service\">";
                echo "<h2>Random on $service</h2>";
                echo "<div class=\"show-list\">";
                $files = array_diff(scandir("/var/episodes/$service/"), array('..', '.'));
                foreach($files as $file) {  
                    echo "<a class=\"show\" href=\"/random/netflix/$file\">";
                    echo "<img src=\"images/$file.jpeg\" alt=\"$file Show Poster\" class=\"show-image\">";
                    echo "<p class=\"show-title\">$file</p>";
                    echo "</a>";
                }
                echo "</div>";
            echo "</div>";
        }
    ?>
  </body>
</html>