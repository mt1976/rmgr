#RMGR
RMGR implements a simple command-line tool for sending and receiving data over a network. The code is organized into several functions and includes several comments that provide context and explanations.

Includes a flag for selecting whether to run in sender or receiver mode, and a flag for specifying the target host and port.

The **send** function is used to send data over the network. It first retrieves the target host and port from the configuration file, and then uses the S.Run function to start the sender process.

The **receive** function is used to receive data over the network. It first retrieves the target host and port from the configuration file, and then uses the R.Run function to start the receiver process.

#### Arguments
Usage of Rmgr is:
**-recv** - runs application in receiver mode
**-send** - runs application in sender mode