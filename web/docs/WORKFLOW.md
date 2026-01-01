# Laravel + Podman Development Workflow

## Initial Setup

### 1. Install Podman
```bash
sudo pacman -S podman podman-compose podman-docker
```

### 2. Create Project Structure
```bash
mkdir ~/laravel-local
cd ~/laravel-local
```

### 3. Create Dockerfile
```dockerfile
FROM php:8.4-fpm

# Install system dependencies
RUN apt-get update && apt-get install -y \
    git curl libpng-dev libonig-dev libxml2-dev zip unzip

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Install PHP extensions
RUN docker-php-ext-install pdo_mysql mbstring exif pcntl bcmath gd

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /var/www
```

### 4. Create docker-compose.yml
```yaml
version: '3.8'

services:
  app:
    image: laravel-local_app
    container_name: laravel-app
    working_dir: /var/www
    volumes:
      - ./:/var/www
    networks:
      - laravel
    depends_on:
      - db

  nginx:
    image: nginx:alpine
    container_name: laravel-nginx
    ports:
      - "8000:80"
    volumes:
      - ./:/var/www
      - ./nginx.conf:/etc/nginx/conf.d/default.conf
    depends_on:
      - app
    networks:
      - laravel

  db:
    image: mariadb:11
    container_name: laravel-db
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: laravel
      MYSQL_USER: laravel
      MYSQL_PASSWORD: secret
    volumes:
      - db-data:/var/lib/mysql
    networks:
      - laravel

networks:
  laravel:
    driver: bridge

volumes:
  db-data:
```

### 5. Create nginx.conf
```nginx
server {
    listen 80;
    index index.php index.html;
    root /var/www/public;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location ~ \.php$ {
        fastcgi_pass app:9000;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }
}
```

---

## Daily Workflow

### Start Development Environment
```bash
# Build the image (first time only)
podman build --no-cache -t laravel-local_app -f Dockerfile .

# Start containers
podman-compose up -d

# Check status
podman ps
```

### Stop Development Environment
```bash
# Stop containers
podman-compose down

# Stop and remove volumes (caution: deletes database)
podman-compose down -v
```

---

## Common Laravel Commands

### Composer
```bash
# Install dependencies
podman exec -it laravel-app composer install

# Update dependencies
podman exec -it laravel-app composer update

# Add package
podman exec -it laravel-app composer require package/name
```

### Artisan
```bash
# Run migrations
podman exec -it laravel-app php artisan migrate

# Create migration
podman exec -it laravel-app php artisan make:migration create_table_name

# Create model with migration and controller
podman exec -it laravel-app php artisan make:model ModelName -mcr

# Create controller
podman exec -it laravel-app php artisan make:controller ControllerName

# Create policy
podman exec -it laravel-app php artisan make:policy PolicyName --model=ModelName

# Clear cache
podman exec -it laravel-app php artisan cache:clear
podman exec -it laravel-app php artisan config:clear
podman exec -it laravel-app php artisan view:clear

# Generate app key
podman exec -it laravel-app php artisan key:generate
```

### NPM/Frontend
```bash
# Install Node dependencies
podman exec -it laravel-app npm install

# Build for production (one-time)
podman exec -it laravel-app npm run build

# Watch for changes (keeps running)
podman exec -it laravel-app npm run dev
```

### Database
```bash
# Access MySQL shell
podman exec -it laravel-db mysql -u laravel -p
# Password: secret

# Seed database
podman exec -it laravel-app php artisan db:seed

# Fresh migration (drops all tables)
podman exec -it laravel-app php artisan migrate:fresh
```

---

## Container Management

### View Logs
```bash
# View all logs
podman-compose logs

# Follow logs
podman-compose logs -f

# Specific service
podman logs laravel-app
podman logs laravel-nginx
podman logs laravel-db
```

### Access Container Shell
```bash
# PHP container
podman exec -it laravel-app bash

# Database container
podman exec -it laravel-db bash

# Nginx container
podman exec -it laravel-nginx sh
```

### Restart Services
```bash
# Restart all
podman-compose restart

# Restart specific service
podman restart laravel-app
podman restart laravel-nginx
podman restart laravel-db
```

---

## Troubleshooting

### Permission Issues
```bash
# Fix Laravel storage permissions
chmod -R 775 storage bootstrap/cache
chmod -R 777 storage/framework
chmod -R 777 storage/logs

# Fix database permissions (SQLite)
chmod 666 database/database.sqlite
chmod 777 database
```

### Rebuild Everything
```bash
# Stop and remove containers
podman-compose down

# Remove image
podman rmi laravel-local_app

# Rebuild from scratch
podman build --no-cache -t laravel-local_app -f Dockerfile .

# Start fresh
podman-compose up -d
```

### View Container Info
```bash
# List images
podman images

# List containers (all)
podman ps -a

# Inspect container
podman inspect laravel-app

# Check Podman info
podman info
```

### Clean Up
```bash
# Remove stopped containers
podman container prune

# Remove unused images
podman image prune -a

# Remove unused volumes
podman volume prune

# Nuclear option: remove everything
podman system prune -a --volumes
```

---

## Environment Configuration

### Update .env for MySQL
```env
DB_CONNECTION=mysql
DB_HOST=db
DB_PORT=3306
DB_DATABASE=laravel
DB_USERNAME=laravel
DB_PASSWORD=secret
```

### Update .env for SQLite
```env
DB_CONNECTION=sqlite
DB_DATABASE=/var/www/database/database.sqlite
```

---

## Access Points

- **Application**: http://localhost:8000
- **Vite Dev Server**: http://localhost:5173 (when running npm run dev)
- **MySQL Port**: localhost:3306 (if exposed in docker-compose.yml)

---

## Quick Reference

```bash
# Start project
cd ~/laravel-local
podman-compose up -d

# Build assets
podman exec -it laravel-app npm run build

# Access app
open http://localhost:8000

# Stop project
podman-compose down
```

---

## Tips

1. **Always run commands inside containers** using `podman exec -it laravel-app`
2. **Keep npm run dev running** in a separate terminal for hot reload
3. **Clear cache** after changing config files or routes
4. **Use MySQL instead of SQLite** in containers for better reliability
5. **Commit your Dockerfile** and docker-compose.yml to version control
6. **Don't commit** .env, vendor/, node_modules/, or storage/ files
