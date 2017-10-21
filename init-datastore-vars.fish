# In bash you can run $(gcloud beta emulators datastore env-init)
set -x DATASTORE_DATASET dev
set -x DATASTORE_EMULATOR_HOST localhost:8081
set -x DATASTORE_EMULATOR_HOST_PATH localhost:8081/datastore
set -x DATASTORE_HOST http://localhost:8081
set -x DATASTORE_PROJECT_ID dev
