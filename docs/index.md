---
page_title: "epic Provider"
subcategory: ""
description: |-
  
---

# Epic Provider



## Example Usage

```terraform
provider "epic" {}

provider "aws" {
  region = "us-east-1"
}

resource "epic_random_name" "movie_name" {
  media_type = "movie"
  title      = "lord of the rings"
  lower      = true
  upper      = false
}


resource "epic_random_quote" "lotr_quote" {
  media_type = "movie"
  title      = "lord of the rings"
}

resource "aws_s3_bucket" "epic" {
  bucket = epic_random_name.movie_name.name

  tags = {
    Name        = epic_random_name.movie_name.name
    Description = epic_random_quote.lotr_quote.quote
  }
}

output "s3_bucket_name" {
  value = aws_s3_bucket.epic.bucket
}
```

## Available Media Types

Available media types are `anime` `movie` `tv_series` and `video_games`

## Media Type: anime Titles:
one_piece
spy_x_family

## Media Type: movie Titles:
jurassic_park
lord_of_the_rings
star_wars

## Media Type: tv_series Titles:
breaking_bad
game_of_thrones

## Media Type: video_game Titles:
final_fantasy_vii
kingdom_hearts_1
the_witcher