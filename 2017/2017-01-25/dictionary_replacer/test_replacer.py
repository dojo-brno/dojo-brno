import pytest
from replacer import replace


def test_empty_template():
    template = ""
    mapping = {}
    expected = ""
    assert replace(template, mapping) == expected

def test_template_without_placeholder():
    template = "no placeholders here"
    mapping = {}
    expected = "no placeholders here"
    assert replace(template, mapping) == expected

def test_empty_mapping_makes_no_replacement():
    template = "$replaceme$"
    mapping = {}
    expected = "$replaceme$"
    assert replace(template, mapping) == expected

def test_mapping():
    template = "$a$"
    mapping = {"a": "b"}
    expected = "b"
    assert replace(template, mapping) == expected

def test_different_mapping():
    template = "$a$"
    mapping = {"a": "c"}
    expected = "c"
    assert replace(template, mapping) == expected

def test_dict_not_empty():
    template = "$b$"
    mapping = {"a":"c"}
    expected = "$b$"
    assert replace(template, mapping) == expected

def test_empty_replacement():
    template = "$b$"
    mapping = {"b":""}
    expected = ""
    assert replace(template, mapping) == expected

def test_empty_replacement2():
    template = "$c$"
    mapping = {"c":""}
    expected = ""
    assert replace(template, mapping) == expected


def test_empty_variable():
    template = "$$"
    mapping = {"": "banana"}
    expected = "banana"
    assert replace(template, mapping) == expected
