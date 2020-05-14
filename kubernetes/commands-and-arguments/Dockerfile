FROM python:3

RUN useradd -ms /bin/bash david
USER david

WORKDIR /users

ADD app.py .

VOLUME /users

# commnad to run when the container starts& override its params with cmd
ENTRYPOINT [ "python", "/users/app.py" ]
CMD [ "hello", "world" ]
