# Simulador de Dispositivos IoT

## 1. Objetivo

Este projeto visa criar um simulador de dispositivos IoT que utiliza o protocolo MQTT para enviar informações simuladas baseadas em dados de sensores reais, nesse caso o [Sensor de Radiação Solar Sem Fio HOBOnet RXW-LIB-900](https://sigmasensors.com.br/produtos/sensor-de-radiacao-solar-sem-fio-hobonet-rxw-lib-900).

## 2. Como Instalar e Rodar

### Pré-requisitos

- Broker MQTT ([Mosquitto](https://mosquitto.org/download/) ou outro de sua preferência)

>[!IMPORTANT]
>Caso queira editar o código ou rodá-lo de outra maneira, tenha instalado [go](https://go.dev/doc/install) em sua máquina.

### Instalação

Clone o repositório para a sua máquina local:

```bash
git clone https://github.com/Lemos1347/inteli-modulo-9-ponderada-1
cd inteli-modulo-9-ponderada-1
```

>[!NOTE]
> Caso queira editar o projeto, lembre de instalar as depencias de cada módulo:

```bash
cd publisher ; go mod tidy ; cd ../subscriber ; go mod tidy ; cd ..
```

### Execução

Para iniciar o simulador, execute primeiro o seu broker, caso seja mosquitto, utilize nossa configuração personalizada:

```bash
mosquitto -c mosquitto.conf
```

E depois em terminais diferentes execute o publisher (lembrando que para o publisher é necessário informar o path correto para o csv contendo os dados simulado do sensor) e o subscriber:

- Publisher:
```bash
./publisher/publisher ./data/dados_sensor_radiacao_solar.csv
```

- Subscriber
```bash
./subscriber/subscriber
```

## 3. Estrutura do Projeto

O projeto é composto por dois módulos go, o publisher e o subscriber (ambos podem ser encontrados em suas respectivas pastas).

- `subscriber/main.go`: Arquivo principal que inicia a parte de subscriber do simulador.
- `publisher/main.go`: Arquivo principal que inicia a parte de publisher do simulador.
- `publisher/solar_sensor.go`: Simula a geração de dados para um sensor de radiação solar.
- `data/dados_sensor_radiacao_solar.csv`: Contém dados simulados para o sensor de radiação solar.
- `mosquitto.conf`: Configuração para o broker MQTT Mosquitto (se aplicável).

## 4. Demonstração do Funcionamento

https://github.com/Lemos1347/inteli-modulo-9-ponderada-1/assets/99190347/9af1a3c8-fc05-4551-bc9c-323ec10e0d7f

