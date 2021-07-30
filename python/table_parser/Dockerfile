FROM python:3.9-slim
WORKDIR /usr/src/app
RUN apt-get update && apt-get install --no-install-recommends -y git libpq-dev && apt-get clean \
 && rm -rf /var/lib/apt/lists/*
COPY requirements.txt .

RUN pip3 install --no-cache-dir -r requirements.txt

COPY . .
RUN python3 setup.py develop

ENTRYPOINT ["table_parser"]