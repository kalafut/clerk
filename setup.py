from setuptools import setup

# pip install --upgrade -e .
setup(
    name='clerk',
    version='0.01',
    py_modules=['clerk'],
    install_requires=[
        'click >= 6.6',
        'python-Levenshtein >= 0.12.0'
    ],
    entry_points='''
        [console_scripts]
        clerk=main:cli
    '''
)
