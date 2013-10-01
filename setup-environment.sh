sudo apt-get update -y -qq
sudo apt-get install git -y -qq

wget https://go.googlecode.com/files/go1.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.1.linux-amd64.tar.gz
export PATH="$PATH:/usr/local/go/bin"
export PATH="$PATH:$GOPATH/bin"

export GOPATH="$HOME/go"
mkdir -p "$GOPATH/{src,bin,pkg}"
mkdir -p "$GOPATH/src/github.com/pjvds/docker-to-digital-ocean-example"
rsync -avz "$WERCKER_SOURCE_DIR/" "$GOPATH/src/github.com/pjvds/docker-to-digital-ocean-example"
export WERCKER_SOURCE_DIR="$GOPATH/src/github.com/pjvds/docker-to-digital-ocean-example"
