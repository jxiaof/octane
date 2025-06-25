import time
import psutil
import yaml

def cpu_performance_test(duration=60):
    print("Starting CPU performance test...")
    start_time = time.time()
    cpu_usage = []
    
    while time.time() - start_time < duration:
        usage = psutil.cpu_percent(interval=1)
        cpu_usage.append(usage)
        print(f"Current CPU Usage: {usage}%")
    
    average_usage = sum(cpu_usage) / len(cpu_usage)
    print(f"Average CPU Usage over {duration} seconds: {average_usage}%")
    
    return {
        'duration': duration,
        'average_usage': average_usage,
        'cpu_usage': cpu_usage
    }

if __name__ == "__main__":
    duration = 60  # Default duration for the test
    results = cpu_performance_test(duration)
    
    # Save results to a YAML file
    with open('cpu_test_results.yaml', 'w') as file:
        yaml.dump(results, file)
    
    print("CPU performance test completed. Results saved to cpu_test_results.yaml.")