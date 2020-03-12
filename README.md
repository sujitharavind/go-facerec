# go-facerec
Requirements -

Go
gocv
OpenCV


Installation steps -

Go: https://golang.org/doc/install

gocv: 
    go get -u -d gocv.io/x/gocv

OpenCV:
    Quick Install
        The following commands should do everything to download and install OpenCV 4.2.0 on Linux:

        cd $GOPATH/src/gocv.io/x/gocv
        make install

If it works correctly, at the end of the entire process, the following message should be displayed:
        
        gocv version: 0.22.0
        opencv lib version: 4.2.0 

Fixes -
Gtk-Message: 18:24:45.381: Failed to load module "canberra-gtk-module"
    sudo apt install libcanberra-gtk-module 