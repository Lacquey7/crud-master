# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  # GATEWAY
  config.vm.define "gateway" do |gateway|
    gateway.vm.box = "hashicorp-education/ubuntu-24-04"
    gateway.vm.hostname = "gateway"
    gateway.vm.network "forwarded_port", guest: 8080, host: 8080, auto_correct: true
    gateway.vm.network "private_network", ip: "192.168.56.10"

    # Chemins résolus
    host_src  = File.expand_path("./srcs/api-gateway")          # dossier à côté du Vagrantfile
    guest_dst = "/home/vagrant/api_gateway_app"                # chemin absolu dans la VM

    # Créer le dossier cible avant copie (en tant que vagrant)
    gateway.vm.provision "shell", privileged: false, inline: <<-SHELL
        mkdir -p #{guest_dst}
    SHELL

    # Copie one-shot (si tu veux juste déposer les fichiers)
    gateway.vm.provision "file", source: host_src + "/", destination: guest_dst

    # Script d’install/launch
    gateway.vm.provision "shell", path: "script/gateway.sh"
  end

  # INVENTORY
  config.vm.define "inventory" do |inventory|
    inventory.vm.box = "hashicorp-education/ubuntu-24-04"
    inventory.vm.hostname = "inventory"
    inventory.vm.network "forwarded_port", guest: 8080, host: 8083, auto_correct: true
    inventory.vm.network "private_network", ip: "192.168.56.11"

    # Chemins résolus
    host_src  = File.expand_path("./srcs/inventory-app")          # dossier à côté du Vagrantfile
    guest_dst = "/home/vagrant/inventory_app"                # chemin absolu dans la VM

    # Créer le dossier cible avant copie (en tant que vagrant)
    inventory.vm.provision "shell", privileged: false, inline: <<-SHELL
       mkdir -p #{guest_dst}
    SHELL

    # Copie one-shot (si tu veux juste déposer les fichiers)
    inventory.vm.provision "file", source: host_src + "/", destination: guest_dst

    # Script d’install/launch
    inventory.vm.provision "shell", path: "script/inventory.sh"

  end

  # BILLING
  config.vm.define "billing" do |billing|
    billing.vm.box = "hashicorp-education/ubuntu-24-04"
    billing.vm.hostname = "billing"

    # Ports (à ajuster selon tes besoins)
    billing.vm.network "forwarded_port", guest: 8080, host: 8082, auto_correct: true
    billing.vm.network "forwarded_port", guest: 5672, host: 5672, auto_correct: true   # RabbitMQ amqp
    billing.vm.network "forwarded_port", guest: 15672, host: 15672, auto_correct: true # RabbitMQ UI (si tu l’utilises)
    billing.vm.network "private_network", ip: "192.168.56.12"

    # Chemins résolus
    host_src  = File.expand_path("./srcs/billing-app")          # dossier à côté du Vagrantfile
    guest_dst = "/home/vagrant/billing_app"                # chemin absolu dans la VM

    # Créer le dossier cible avant copie (en tant que vagrant)
    billing.vm.provision "shell", privileged: false, inline: <<-SHELL
      mkdir -p #{guest_dst}
    SHELL

    # Copie one-shot (si tu veux juste déposer les fichiers)
    billing.vm.provision "file", source: host_src + "/", destination: guest_dst

    # Script d’install/launch
    billing.vm.provision "shell", path: "script/billing.sh"
  end
end