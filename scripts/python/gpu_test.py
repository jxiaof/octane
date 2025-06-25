import pynvml
import time
import torch

def initialize_gpu():
    pynvml.nvmlInit()
    device_count = pynvml.nvmlDeviceGetCount()
    return device_count

def get_gpu_info(device_index):
    handle = pynvml.nvmlDeviceGetHandleByIndex(device_index)
    info = {
        'name': pynvml.nvmlDeviceGetName(handle).decode('utf-8'),
        'memory_total': pynvml.nvmlDeviceGetMemoryInfo(handle).total,
        'memory_free': pynvml.nvmlDeviceGetMemoryInfo(handle).free,
        'memory_used': pynvml.nvmlDeviceGetMemoryInfo(handle).used,
        'temperature': pynvml.nvmlDeviceGetTemperature(handle, pynvml.NVML_TEMPERATURE_GPU),
        'utilization': pynvml.nvmlDeviceGetUtilizationRates(handle).gpu,
    }
    return info

def run_tensor_operations(device_index):
    device = torch.device(f'cuda:{device_index}')
    tensor_size = 10000
    a = torch.randn(tensor_size, device=device)
    b = torch.randn(tensor_size, device=device)
    
    start_time = time.time()
    c = torch.matmul(a, b)
    end_time = time.time()
    
    return end_time - start_time

def main():
    device_count = initialize_gpu()
    print(f"Detected {device_count} GPU(s).")
    
    for i in range(device_count):
        gpu_info = get_gpu_info(i)
        print(f"GPU {i}: {gpu_info}")
        
        duration = run_tensor_operations(i)
        print(f"Tensor operations on GPU {i} took {duration:.4f} seconds.")

    pynvml.nvmlShutdown()

if __name__ == "__main__":
    main()