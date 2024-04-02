echo "Step 1"
go run main.go ./data/step1/invalid.json
go run main.go ./data/step1/valid.json

echo "Step 2"
go run main.go ./data/step2/invalid.json 
go run main.go ./data/step2/invalid2.json
go run main.go ./data/step2/valid.json
go run main.go ./data/step2/valid2.json

echo "Step 3"
go run main.go ./data/step3/invalid.json
go run main.go ./data/step3/valid.json

echo "Step 4"
go run main.go ./data/step4/invalid.json
go run main.go ./data/step4/valid.json
go run main.go ./data/step4/valid2.json