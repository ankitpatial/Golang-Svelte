import gzip
import json
import os
import tempfile
import uuid

import boto3
from weasyprint import HTML


class Response:
    def __init__(self, bucket, key, folder, dir, section, fileName, contentType, contentSize):
        self.bucket = bucket
        self.key = key
        self.folder = folder
        self.dir = dir
        self.section = section
        self.fileName = fileName
        self.contentType = contentType
        self.contentSize = contentSize


def lambda_handler(event, context):
    payload = event["Payload"]
    dest = payload["dest"]

    uid = uuid.uuid5(uuid.NAMESPACE_DNS, "rfx.services")
    name = "{}.html".format(uid)
    html_path = os.path.join(tempfile.gettempdir(), name)

    print("going to save html ", html_path)
    f = open(html_path, "a")
    f.write(payload["html"])
    f.close()
    print("save html at", html_path)

    # Call the WeasyPrint command
    pdf_path = os.path.join(tempfile.gettempdir(), "{}.pdf".format(uid))
    HTML(html_path).write_pdf(pdf_path)

    # upload pdf to destination
    bucket = payload["bucket"]
    key = payload["key"]
    upload_file_to_s3(pdf_path, bucket, key)
    size = os.path.getsize(pdf_path)

    # remove local files
    remove_file(html_path)
    remove_file(pdf_path)

    res = Response(bucket, key, dest["folder"], dest["dir"], dest["section"], dest["fileName"], "application/pdf", size)
    return json.dumps(res.__dict__)


def upload_file_to_s3(filepath, bucket, key):
    s3_client = boto3.client('s3')
    try:
        f = open(filepath, "rb")
        data = f.read()
        f.close()

        data_gz = gzip.compress(data)
        s3_client.put_object(
            Bucket=bucket, Key=key, Body=data_gz, ContentEncoding="gzip", ContentType="application/pdf"
        )
        print(f"File '{filepath}' uploaded to '{bucket}/{key}' successfully.")
        return True
    except FileNotFoundError:
        print(f"File '{filepath}' not found.")
    except Exception as e:
        print(f"An error occurred while uploading file '{filepath}': {str(e)}")

    return False


def remove_file(path):
    try:
        os.remove(path)
        print(f"File '{path}' removed successfully.")
    except FileNotFoundError:
        print(f"File '{path}' not found.")
    except PermissionError:
        print(f"Permission denied to remove file '{path}'.")
    except Exception as e:
        print(f"An error occurred while removing file '{path}': {str(e)}")
