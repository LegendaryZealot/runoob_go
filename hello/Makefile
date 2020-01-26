TARGET  := app

.PHONY : everything objs clean veryclean rebuild

everything : build

all : build

run:build
	./$(TARGET)
build : 
	go build -o $(TARGET) main.go
rebuild: veryclean everything
                
clean :
	-rm -f $(OBJS)
    
veryclean : clean
	-rm -f $(TARGET)