terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
    }
  }
}

provider "docker" {}

resource "docker_image" "mysql" {
  name         = "mysql:8"
}

resource "random_password" "mysql_root_password" {
  length = 16
}

resource "docker_container" "mysql" {
  name  = "mysql"
  image = "${docker_image.mysql.image_id}"
  env {
    MYSQL_ROOT_PASSWORD = "${random_password.mysql_root_password}"
  }
  mounts {
    target = "./mounts/mysql/data"
    source = "/var/lib/mysql/data"
    type = "bind"
  }
  ports {
    internal = 3306
    external = 3306
  }
}


#provider "mysql" {
#  endpoint = "127.0.0.1:3306"
#  username = "root"
#  password = "${random_password.mysql_root_password}"
#}
#resource "mysql_database" "test_db" {
#  name = "test_db"
#}