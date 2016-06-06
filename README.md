hanson
======

___!!! HANSON IS EXTREMLY WORK IN PROGRESS !!!___

A simple Ping-Server for websites. Integrate a tiny JavaScript to your community-site for watching your currently users, live.

Building
--------
You have to install some prerequisites for building hanson for productive server. We have to transpile all the reactjs code. So you have to install webpack for building.

Install webpack via npm:

```
npm i html-webpack-plugin --save-dev
```


The view components uses ReactJS. Before starting hanson you have to build the whole project. We created a Makefile with the following targets:

|      Targets       | Description                                                                           |
| ------------------ | ------------------------------------------------------------------------------------- |
| `make build`       | Builds the whole project with precompiling all CSS and bundle the ReactJS components. |
| `make buildrun`    | Builds the whole project (see build) and runs the built application from build dir.   |
| `make run`         | Starts the prebuilt application from build directory.                                 |


Config
------
Before you start hanson you have to set all parameters in config.yaml. Please set the following parameters:

| Parameter          | Description                                                            | Default      |
| ------------------ | ---------------------------------------------------------------------- | ------------ |
| LoggingFileUse     | true/false - Logging on stdout (false) or in a log-file (true).        | false        |
| LoggingFile        | filepath and filename for log-file.                                  | log/hanson.log |
| LoggingLevel       | Detail level of logging. (debug/info/warning/error/fatal/panic)        | warning      |
| WebserverDebug     | Webserver-mode for debug informations (true/false)                     | false        |
| WebserverPort      | Webserver listens on given port                                        | 8080         |


The file config.yaml will be searched at this directories:

1. /etc/hanson/
2. $HOME/.hanson
3. .  -  [hanson application directory]
