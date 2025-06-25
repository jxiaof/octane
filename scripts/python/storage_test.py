import os
import time
import psutil
import yaml

def storage_test(path='/'):
    # Get initial disk usage
    initial_usage = psutil.disk_usage(path)
    print(f"Initial disk usage for {path}: {initial_usage.percent}% used")

    # Start the test
    print("Starting storage performance test...")
    start_time = time.time()

    # Create a temporary file to test write speed
    test_file = os.path.join(path, 'storage_test_file.tmp')
    with open(test_file, 'wb') as f:
        f.write(os.urandom(1024 * 1024 * 100))  # Write 100 MB of random data

    write_time = time.time() - start_time
    print(f"Write test completed in {write_time:.2f} seconds")

    # Measure read speed
    start_time = time.time()
    with open(test_file, 'rb') as f:
        f.read()  # Read the file

    read_time = time.time() - start_time
    print(f"Read test completed in {read_time:.2f} seconds")

    # Clean up the test file
    os.remove(test_file)

    # Get final disk usage
    final_usage = psutil.disk_usage(path)
    print(f"Final disk usage for {path}: {final_usage.percent}% used")

    # Prepare results
    results = {
        'initial_usage': initial_usage.percent,
        'final_usage': final_usage.percent,
        'write_time': write_time,
        'read_time': read_time
    }

    return results

if __name__ == "__main__":
    results = storage_test()
    with open('storage_test_results.yaml', 'w') as yaml_file:
        yaml.dump(results, yaml_file)
    print("Storage test results saved to storage_test_results.yaml")