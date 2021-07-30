from setuptools import setup

setup(
    name='table_parser',
    entry_points={
        'console_scripts': [
            'table_parser=table_parser.cmd.parser:main'
        ],
    }

)
