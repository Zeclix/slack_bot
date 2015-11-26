import posixpath as path
from fabric.api import run
from fabric.context_managers import cd
from fabric.contrib.files import exists
from fabric.operations import put

COPY_FILES = (
    'slack_bot'
)

def deploy(path):
    with cd(path):
        for copy_file in COPY_FILES:
            if exists(copy_file):
                run('rm -f %s' % copy_file)
            put(copy_file, path)
        if exists('bot.pid'):
            pid = run('cat bot.pid')
            run('kill -9 ' + pid, quiet=True)
        run('./slack_bot')
