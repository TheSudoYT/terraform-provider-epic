provider "epic" {}

provider "aws" {
  region = "us-east-1"
}

resource "epic_random_name" "movie_name" {
  media_type = "movie"
  title      = "lord of the rings"
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