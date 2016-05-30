import config


def test_config():
    cfg = config.load('config.json')
    assert cfg['import']['sources']['chase_checking']['format'] == 'csv'
