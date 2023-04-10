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
          $shows = array_diff(scandir("/var/episodes/$service/"), array('..', '.'));
          foreach($shows as $show) {  
            $jsonString=file_get_contents("/var/episodes/$service/$show/metadata.json");
            $jsonData=json_decode($jsonString,true);

            $slug = $jsonData["slug"];
            $imgRef = $jsonData["imgRef"];
            
            echo "<a class=\"show\" href=\"/random/$service/$show\">";
            echo "<img src=\"$imgRef\" alt=\"$slug Show Poster\" class=\"show-image\">";
            echo "<p class=\"show-title\">$slug</p>";
            echo "</a>";
          }
          echo "</div>";
        echo "</div>";
      }
    ?>
  </body>
</html>