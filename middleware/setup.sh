echo "Copying Connection Profiles"
cp ../connections/connection*.json ./fabric/ccp/

echo "Starting MongoDB Container"
docker run --name healthmongo -d -p 27017:27017 mongo:latest

echo "Running npm install"
npm install

echo "Enrolling Admins"
node ./setup/enroll_admin.js
