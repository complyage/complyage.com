# Use an official PHP image with Apache
FROM php:8.2-apache

# Copy your PHP files into the container
COPY . /var/www/html/

RUN a2enmod rewrite


# Expose port 80 (Apache default)
EXPOSE 80
