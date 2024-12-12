variable "public_key" {
  type        = string
  description = "Path to the public ssh key"
  default     = "/root/.ssh/id_rsa.pub"
}
