import subprocess
import json
import time
import psutil

def run_speedtest():
    """Run a speed test and return the results."""
    try:
        result = subprocess.run(['speedtest', '--json'], capture_output=True, text=True, check=True)
        return json.loads(result.stdout)
    except Exception as e:
        print(f"Error running speed test: {e}")
        return None

def get_network_info():
    """Collect network interface information."""
    interfaces = psutil.net_if_addrs()
    network_info = {}
    
    for interface, addrs in interfaces.items():
        network_info[interface] = {
            'addresses': [addr.address for addr in addrs if addr.family == psutil.AF_INET],
            'mac': [addr.address for addr in addrs if addr.family == psutil.AF_LINK]
        }
    
    return network_info

def main():
    """Main function to run network tests."""
    print("Collecting network information...")
    network_info = get_network_info()
    print("Network Information:", network_info)

    print("Running speed test...")
    speedtest_results = run_speedtest()
    
    if speedtest_results:
        print("Speed Test Results:", speedtest_results)
    else:
        print("Failed to retrieve speed test results.")

if __name__ == "__main__":
    main()