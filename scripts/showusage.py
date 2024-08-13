import json, psutil, datetime, time


def get_pid_by_name(process_name):
    for proc in psutil.process_iter(['pid', 'name']):
        if proc.info['name'] == process_name:
            return proc.info['pid']
    return None

while True:
    try:
        pid = get_pid_by_name("top10url")

        if pid is not None:

            print(json.dumps((datetime.datetime.now().isoformat(), 
                            "rss(Mb)", psutil.Process(int(pid)).memory_info()._asdict()["rss"]/ (1024*1024),
                            "cpu(%)", psutil.Process(pid).cpu_percent(interval=1)))
                )
        else:
            print("Process not found")

    except Exception as e:
        print(f"ERROR: {e}")

    time.sleep(1)
