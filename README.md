# RMGR

RMGR implements a simple command-line tool for sending and receiving data over a network. The code is organized into several functions and includes comments providing context and explanations.

It includes a flag for selecting whether to run in sender or receiver mode and a flag for specifying the target host and port.

The **send** function sends data over the network. It first retrieves the target host and port from the configuration file and then uses the S.Run function to start the sender process.

The **receive** function is used to receive data over the network. It first retrieves the target host and port from the configuration file, and then uses the R.Run function to start the receiver process.

## Arguments
Usage of Rmgr is:

**-recv** - runs the application in receiver mode

**-send** - runs the application in sender mode
