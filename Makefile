# build-folder
BUILD_FOLDER := build

# application-name
APPNAME := hanson



# start pre-builded application
run:
				cd $(BUILD_FOLDER); ./$(APPNAME)


# Build application and runs webserver
buildrun:	build run


# builds application
build: 	clean csscompile copyelements reactcompile
				go build -o $(BUILD_FOLDER)/$(APPNAME)


# create build folder and copies all necessary files
copyelements:
				mkdir -p $(BUILD_FOLDER)
				cp -R public $(BUILD_FOLDER)/public
				rm -Rf $(BUILD_FOLDER)/public/gcss
				rm -Rf $(BUILD_FOLDER)/public/js/*
				cp -R views $(BUILD_FOLDER)/views
				mkdir -p $(BUILD_FOLDER)/log
				cp config.yaml $(BUILD_FOLDER)


# compiles all ReactJS Components
reactcompile:
				./node_modules/.bin/webpack

# precompile all gcss to css files
csscompile:
				go run main.go -precompileonly


# cleans project-build-data
clean:
				rm -rf $(BUILD_FOLDER)
