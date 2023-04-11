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
        echo '<div class="service">';
          echo "<h2>Random on $service</h2>";
          echo '<div class="show-list">';
          $shows = array_diff(scandir("/var/episodes/$service/"), array('..', '.'));
          foreach($shows as $show) {
            $jsonString=file_get_contents("/var/episodes/$service/$show/metadata.json");
            $jsonData=json_decode($jsonString,true);

            $slug = $jsonData["slug"];
            $imgRef = $jsonData["imgRef"];
            $numSeasons = $jsonData["numSeasons"];
            
            echo '<div class="show">';
              echo "<a class=\"show-all\" href=\"/random/$service/$show\">";
                echo "<img src=\"$imgRef\" alt=\"$slug Show Poster\" class=\"show-image\">";
                echo "<p class=\"show-title\">$slug</p>";
              echo '</a>';
              echo "<input type=\"checkbox\" id=\"switch-$show\" class=\"toggle\">";
              echo "<label class=\"lbl-toggle\" for=\"switch-$show\">By Seasons</label>";
              echo '<ul class="season-list collapsible-content">';
              for ($season = 1; $season <= $numSeasons; $season++) {
                echo '<li class="season">';
                  echo "<a href=\"/random/$service/$show/s$season\">Season $season</a>";
                echo '</li>';
              }
              echo '</input>';
            echo '</div>';
          }
          echo '</div>';
        echo '</div>';
      }
    ?>
  </body>
</html>