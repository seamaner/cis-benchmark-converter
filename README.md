# cis-benchmark-to-csv
Converts CIS Benchmark PDFs into usable CSV & Excel files.

For expected numbers see the [reference results](reference_results.md) of the conversions.

## Requirements
- Python 3
- For Excel output install the `xlsxwriter` Python package

## Dumping Text
```
Complie pdf2txt.go to pdf2txt then:
- pdf2txt: ./pdf2txt CIS_Kubernetes_Benchmark_v1.6.0.pdf > CIS_Kubernetes_Benchmark_v1.6.0.txt
- Convert txt to excel:  python3 cis2excel.py CIS_Kubernetes_Benchmark_v1.6.0.txt
- Convert xlsx to yaml
```

## Program Usage
```
usage: cis2csv.py [-h] [-l {DEBUG,INFO,WARNING,ERROR,CRITICAL}] inputFilePath

positional arguments:
  inputFilePath         CIS text dump to parse.

optional arguments:
  -h, --help            show this help message and exit
  -l {DEBUG,INFO,WARNING,ERROR,CRITICAL}, --log-level {DEBUG,INFO,WARNING,ERROR,CRITICAL}
                        Set the logging level
```

```
usage: cis2excel.py [-h] [-l {DEBUG,INFO,WARNING,ERROR,CRITICAL}]
                    [--sheetName SHEETNAME]
                    inputFilePath

positional arguments:
  inputFilePath         CIS text dump to parse.

optional arguments:
  -h, --help            show this help message and exit
  -l {DEBUG,INFO,WARNING,ERROR,CRITICAL}, --log-level {DEBUG,INFO,WARNING,ERROR,CRITICAL}
                        Set the logging level
  --sheetName SHEETNAME
                        Set the sheet name in Excel.
```
