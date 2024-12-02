:: Set variables for image name and tag
set IMAGE_NAME=grand-oak-service
set IMAGE_TAG=2.0
set DOCKERFILE_PATH=src/grand-oak-service
set USER_NAME=dhanikanovlisa

:: Print a message
echo Building Docker image: %IMAGE_NAME%:%IMAGE_TAG%

:: Build the Docker image
docker build -t %IMAGE_NAME%:%IMAGE_TAG% %DOCKERFILE_PATH%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to build Docker image.
    exit /b 1
)

:: Login to Docker
docker login
if %ERRORLEVEL% NEQ 0 (
    echo Docker login failed.
    exit /b 1
)

:: Print a message
echo Tagging Docker image: %IMAGE_NAME%:%IMAGE_TAG%

:: Tag the Docker image
docker tag %IMAGE_NAME%:%IMAGE_TAG% %USER_NAME%/%IMAGE_NAME%:%IMAGE_TAG%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to tag Docker image.
    exit /b 1
)

:: Push the Docker image
docker push %USER_NAME%/%IMAGE_NAME%:%IMAGE_TAG%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to push Docker image.
    exit /b 1
)

:: Print success message
echo Docker image %IMAGE_NAME%:%IMAGE_TAG% built and pushed successfully!


:: Set variables for image name and tag
set IMAGE_NAME=pine-valley-service-go
set IMAGE_TAG=2.0
set DOCKERFILE_PATH=src/pine-valley-service
set USER_NAME=dhanikanovlisa

:: Print a message
echo Building Docker image: %IMAGE_NAME%:%IMAGE_TAG%

:: Build the Docker image
docker build -t %IMAGE_NAME%:%IMAGE_TAG% %DOCKERFILE_PATH%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to build Docker image.
    exit /b 1
)


:: Print a message
echo Tagging Docker image: %IMAGE_NAME%:%IMAGE_TAG%

:: Tag the Docker image
docker tag %IMAGE_NAME%:%IMAGE_TAG% %USER_NAME%/%IMAGE_NAME%:%IMAGE_TAG%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to tag Docker image.
    exit /b 1
)

:: Push the Docker image
docker push %USER_NAME%/%IMAGE_NAME%:%IMAGE_TAG%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to push Docker image.
    exit /b 1
)

:: Print success message
echo Docker image %IMAGE_NAME%:%IMAGE_TAG% built and pushed successfully!


:: Set variables for image name and tag
set IMAGE_NAME=doctor-booking
set IMAGE_TAG=2.0
set DOCKERFILE_PATH=src/doctor-booking
set USER_NAME=dhanikanovlisa

:: Print a message
echo Building Docker image: %IMAGE_NAME%:%IMAGE_TAG%

:: Build the Docker image
docker build -t %IMAGE_NAME%:%IMAGE_TAG% %DOCKERFILE_PATH%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to build Docker image.
    exit /b 1
)

:: Login to Docker
docker login
if %ERRORLEVEL% NEQ 0 (
    echo Docker login failed.
    exit /b 1
)

:: Print a message
echo Tagging Docker image: %IMAGE_NAME%:%IMAGE_TAG%

:: Tag the Docker image
docker tag %IMAGE_NAME%:%IMAGE_TAG% %USER_NAME%/%IMAGE_NAME%:%IMAGE_TAG%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to tag Docker image.
    exit /b 1
)

:: Push the Docker image
docker push %USER_NAME%/%IMAGE_NAME%:%IMAGE_TAG%
if %ERRORLEVEL% NEQ 0 (
    echo Failed to push Docker image.
    exit /b 1
)

:: Print success message
echo Docker image %IMAGE_NAME%:%IMAGE_TAG% built and pushed successfully!
exit /b 0