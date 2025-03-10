import re
import matplotlib.pyplot as plt

# Sample log file name (change as needed)
TRACE_FILE = "../trace.txt"

SCHED_REGEX = re.compile(
    r"SCHED (\d+)ms: .*?idleprocs=(\d+).*?runqueue=(\d+) \[([\d\s]+)\] .*?schedticks=\[([\d\s]+)\]"
)

def parse_trace_line(line):
    match = SCHED_REGEX.search(line)
    if not match:
        return None

    time_ms = int(match.group(1))
    idle_procs = int(match.group(2))
    runqueue = int(match.group(3))

    # Parse runqueue array
    runqueue_array = list(map(int, match.group(4).split()))

    # Parse schedticks array
    schedticks = list(map(int, match.group(5).split()))

    return {
        "time_s": time_ms / 1000.0,  # Convert ms to seconds
        "idle_procs": idle_procs,
        "runqueue": runqueue,
        "runqueue_array": runqueue_array,  # Sum of runqueue array
        "schedticks": schedticks,  # Sum of schedticks array
    }


# Read and parse the trace file
def read_trace_file(filename):
    events = []
    with open(filename, "r") as file:
        for line in file:
            event = parse_trace_line(line)
            if event:
                events.append(event)
    return events

# Function to plot the extracted metrics
def plot_metrics(events):
    times = [e["time_s"] for e in events]
    runqueues = [e["runqueue"] for e in events]
    idle_procs = [e["idle_procs"] for e in events]
    # schedticks_sums = [e["schedticks"][0] for e in events]
    runqueue_array = [e["runqueue_array"][0] for e in events]
    runqueue_array_second = [e["runqueue_array"][1] for e in events]
    schedticks_second = [e["schedticks"][1] for e in events]
    runqueue_array_third = [e["runqueue_array"][2] for e in events]
    runqueue_array_fourth = [e["runqueue_array"][3] for e in events]
    




    plt.figure(figsize=(10, 5))
    
    plt.plot(times, runqueues, label="Global Run Queue")
    # plt.plot(times, idle_procs, label="Idle Procs")
    # plt.plot(times, schedticks_sums, label="Sched Ticks 1")
    plt.plot(times, runqueue_array, label="Local Run Queue 1")
    plt.plot(times, runqueue_array_second, label="Local Run Queue 2")
    plt.plot(times, runqueue_array_third, label="Local Run Queue 3")
    plt.plot(times, runqueue_array_fourth, label="Local Run Queue 4")

    # plt.plot(times, schedticks_second, label="Sched Ticks 2")




    plt.xlabel("Time (s)")
    plt.ylabel("Value")
    plt.title("Go Scheduler Trace Analysis")
    plt.legend()
    plt.grid(True)
    
    plt.savefig("sched_trace.png")
    plt.show()

# Main execution
if __name__ == "__main__":
    events = read_trace_file(TRACE_FILE)
    if events:
        plot_metrics(events)
        print("Plot saved as sched_trace.png")
    else:
        print("No valid trace data found.")
