import psutil
import platform
import json
import yaml

def get_system_info():
    system_info = {}

    # 获取操作系统信息
    system_info['os'] = platform.system()
    system_info['os_version'] = platform.version()
    system_info['architecture'] = platform.architecture()[0]
    system_info['hostname'] = platform.node()
    system_info['uptime'] = get_uptime()

    # 获取CPU信息
    system_info['cpu'] = {
        'physical_cores': psutil.cpu_count(logical=False),
        'total_cores': psutil.cpu_count(logical=True),
        'frequency': psutil.cpu_freq()._asdict(),
        'usage': psutil.cpu_percent(interval=1)
    }

    # 获取内存信息
    virtual_memory = psutil.virtual_memory()
    system_info['memory'] = {
        'total': virtual_memory.total,
        'available': virtual_memory.available,
        'used': virtual_memory.used,
        'percent': virtual_memory.percent
    }

    # 获取磁盘信息
    disk_usage = psutil.disk_usage('/')
    system_info['disk'] = {
        'total': disk_usage.total,
        'used': disk_usage.used,
        'free': disk_usage.free,
        'percent': disk_usage.percent
    }

    # 获取网络信息
    system_info['network'] = psutil.net_if_addrs()

    return system_info

def get_uptime():
    return int(psutil.boot_time())

def save_to_yaml(data, filename='system_info.yaml'):
    with open(filename, 'w') as file:
        yaml.dump(data, file)

def save_to_json(data, filename='system_info.json'):
    with open(filename, 'w') as file:
        json.dump(data, file, indent=4)

if __name__ == "__main__":
    info = get_system_info()
    save_to_yaml(info)
    save_to_json(info)