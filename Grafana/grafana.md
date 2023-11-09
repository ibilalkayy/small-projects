# **Grafana Commands**

These are some of the commands that you need to run in order to run Grafana, Loki and Promtail.

## **Install Grafana on Debian and Ubuntu**

    sudo apt-get install -y apt-transport-https
    sudo apt-get install -y software-properties-common wget
    sudo wget -q -O /usr/share/keyrings/grafana.key https://apt.grafana.com/gpg.key

## **Stable Release**

    echo "deb [signed-by=/usr/share/keyrings/grafana.key] https://apt.grafana.com stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list

## **Update the Machine**

    sudo apt-get update

## **Install Grafana**

    sudo apt-get install grafana

## **Start, Enable and check the Status of Grafana**

    sudo /bin/systemctl start grafana-server
    sudo /bin/systemctl enable grafana-server
    sudo /bin/systemctl status grafana-server

# **Install Loki and Promtail using Docker**

## **Download Loki Config**

    wget https://raw.githubusercontent.com/grafana/loki/v2.8.0/cmd/loki/loki-local-config.yaml -O loki-config.yaml

## **Download Promtail Config**

    wget https://raw.githubusercontent.com/grafana/loki/v2.8.0/clients/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml

## **Run Loki Docker container**

    docker run -d --name loki -v $(pwd):/mnt/config -p 3100:3100 grafana/loki:2.8.0 --config.file=/mnt/config/loki-config.yaml

## **Run Promtail Docker container**

    docker run -d --name promtail -v $(pwd):/mnt/config -v /var/log:/var/log --link loki grafana/promtail:2.8.0 --config.file=/mnt/config/promtail-config.yaml