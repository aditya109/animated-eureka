# /usr/bin/bash
swagger_location=$(which swagger)

function execute_swagger() {
	swagger generate spec -t=. -o ./api/swagger/swagger.yaml --scan-models
}

if [[ -z "$swagger_location" ]]; then
	download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
	echo $download_url
	echo "❌ swagger not found..."
	echo "🛒 installing swagger binary..."
	# sudo curl -o /usr/local/bin/swagger -L'#' "$download_url"
	# sudo chmod +x /usr/local/bin/swagger
fi

echo "👻 generating swagger specs with scan-models enabled..."
echo "swagger generate spec -t=. -o ./api/swagger/swagger.yaml --scan-models"

execute_swagger
