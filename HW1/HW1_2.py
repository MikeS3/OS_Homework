#Chatgpt help restructuring and implement threads
import threading
import time
import os
import sys

# Function for the producer thread, writes numbers to the pipe
def producer(pipe_writer, event):
    for i in range(1, 6):
        pipe_writer.write(f"{i}\n")  # Write number as a string followed by a newline
        pipe_writer.flush()  # Ensure the data is sent immediately
        print("Producer:", i)  # Print the produced number
        time.sleep(0.5)  # Simulate work by sleeping for 500ms
    pipe_writer.close()  # Close the write end of the pipe when done
    event.set()  # Signal the consumer that production is complete

# Function for the consumer thread, reads numbers from the pipe
def consumer(pipe_reader, event):
    while not event.is_set() or not pipe_reader.closed:
        line = pipe_reader.readline()  # Read a line from the pipe
        if not line:
            break  # Stop reading if nothing is received
        print("Consumer:", line.strip())  # Print the consumed number
    pipe_reader.close()  # Close the read end of the pipe

# Main function that sets up the producer and consumer threads
def main():
    r, w = os.pipe()  # make the pipe
    pipe_reader = os.fdopen(r, "r")  # Open read end of the pipe
    pipe_writer = os.fdopen(w, "w")  # Open write end of the pipe

    event = threading.Event()  # created an event to signal when complete

    # the producer and consumer threads
    producer_thread = threading.Thread(target=producer, args=(pipe_writer, event))
    consumer_thread = threading.Thread(target=consumer, args=(pipe_reader, event))

    producer_thread.start()  # Start the producer thread
    consumer_thread.start()  # Start the consumer thread

    producer_thread.join()  # Wait for the producer thread to finish
    consumer_thread.join()  # Wait for the consumer thread to finish

if __name__ == "__main__":
    main()

