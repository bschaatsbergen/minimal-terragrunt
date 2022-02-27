resource "aws_iam_role" "default" {
  name = "example-lambda-role-${var.environment}"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Action": "sts:AssumeRole"
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy" "default" {
  name = "example-lambda-role-policy-${var.environment}"
  role = aws_iam_role.default.id

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ReadOnlyBucketAccess",
      "Action": [
        "s3:GetObject",
        "s3:ListBucket"
      ],
      "Effect": "Allow",
      "Resource": [
        "${var.default_bucket_arn}",
        "${var.default_bucket_arn}/*"
      ]
    }
  ]
}
POLICY
}
