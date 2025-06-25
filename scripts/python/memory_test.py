import psutil
import time
import yaml

def memory_test(duration=60):
    start_time = time.time()
    memory_usage = []

    while time.time() - start_time < duration:
        mem = psutil.virtual_memory()
        memory_usage.append({
            'timestamp': time.time(),
            'used_memory': mem.used,
            'available_memory': mem.available,
            'total_memory': mem.total,
            'memory_percent': mem.percent
        })
        time.sleep(1)

    return memory_usage

def save_results_to_yaml(results, filename='memory_test_results.yaml'):
    with open(filename, 'w') as file:
        yaml.dump(results, file)

if __name__ == "__main__":
    duration = 60  # Test duration in seconds
    results = memory_test(duration)
    save_results_to_yaml(results)