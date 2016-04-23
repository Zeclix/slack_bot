import posixpath as path
from fabric.api import run
from fabric.context_managers import cd
from fabric.contrib.files import exists
from fabric.operations import put, local

COPY_FILES = [
    'slack_bot',
    'config'
]

def decrypt_config(name):
    local('rm -f config')
    local('openssl enc -aes-256-ecb -d -in %s -out config -pass env:SLACK_CONFIG_PASS' % name)

def deploy(path):
    with cd(path):
        for copy_file in COPY_FILES:
            if exists(copy_file):
                run('rm -f %s' % copy_file)
            put(copy_file, path, mode=0755)
        if exists('bot.pid'):
            pid = run('cat bot.pid')
            run('kill -9 ' + pid, quiet=True)
        run('./slack_bot')
