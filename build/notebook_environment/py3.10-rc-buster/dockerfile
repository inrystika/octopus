FROM python:3.10-rc-buster
RUN pip install jupyterlab
CMD ["sh","-c", "jupyter lab --no-browser --ip=0.0.0.0 --allow-root --notebook-dir='/code' --port=8888 --LabApp.token='' --LabApp.allow_origin='*' --LabApp.base_url=$OCTOPUS_JPY_BASE_URL"]