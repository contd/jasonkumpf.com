---
title: Today I Learned Python
lang: en-US
tags:
  - python
---

# {{ $page.title }}

## Mutable default arguments

__NOTE__: Information taken from [python guide](http://docs.python-guide.org/en/latest/writing/gotchas/)

One of the most confusing moments for new developers is when they discover how Python treats default arguments in function definitions.

Let's say you want to define a function that accepts a list as a parameter, and you want the default value of that list to be the empty list, so you write:

```python
def append_to(element, to=[]):
    to.append(element)
    return to
```

If everything was okay, you would expect the following behavior:

```python
my_list = append_to(12)
print my_list

my_other_list = append_to(42)
print my_other_list
```

```
[12]
[42]
```

But what you get instead is:

```
[12]
[12, 42]
```

A new list is created once when the function is defined, and the same list is used in each successive call.

Python’s default arguments are evaluated once when the function is defined, instead of each time the function is called (like it is in say Ruby). This means that if you use a mutable default argument and mutate it, you will and have mutated that object for all future calls to the function as well.

### What you should do

Create a new object each time the function is called, by using a default arg to signal that no argument was provided (`None` is often a good choice):

```python
def append_to(element, to=None):
    if to is None:
        to = []
    to.append(element)
    return to
```

## Setuptools entry points

When building a Python package, the most common way of exposing a command line interface (CLI) is to use the "scripts" keyword in `setup.py`.

```python
# file: setup.py
from setuptools import setup

setup(
  name='my_package',
  scripts=['scripts/my_cli'],
  ...
)
```

However, **entry points** represents an improved and cross-platform approach to accomplish the same goal.

```python
# file: setup.py
from setuptools import setup

setup(
  name='my_package',
  entry_points={
    'console_scripts': [
      'my_cli': 'my_package.cli:main'
    ]
  },
  ...
)
```

What's happens when the user installs the package is that ``pip`` will make sure that ``my_cli`` is available as an executable on the system. Once invoked, the ``main`` function of the ``my_package.cli`` module
will be called.

```bash
pip install my_package
my_cli  # the main function is called
Usage: my_cli [OPTIONS] COMMAND [ARGS]...
```

The problem with the former approach is that it requires separating CLI related code into a isolated file outside the main package. With **entry points** it's possible to unify all source code under the same package directory. This makes testing significantly easier.

---

There's actually a lot more you can use entry points for. It is ideally suited to act as a simple way for third party plugins to hook into your application. A great example of such integration is employed by `pytest` for their excellent plugin architecture.

You can also find out more in the `setuptools` docs.


[pytest]: http://pytest.org/latest/
[setuptools]: https://pythonhosted.org/setuptools/setuptools.html#automatic-script-creation
