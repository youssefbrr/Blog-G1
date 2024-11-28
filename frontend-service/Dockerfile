# Use an official Node.js image
FROM node:18

# Install pnpm globally
RUN npm install -g pnpm

# Set the working directory in the container
WORKDIR /app

# Copy package.json, pnpm-lock.yaml, and .npmrc if present
COPY package.json pnpm-lock.yaml* .npmrc* ./

# Install dependencies using pnpm
RUN pnpm install --frozen-lockfile

# Copy the rest of the application code
COPY . .


# Expose the port your application will run on
EXPOSE 3000

# Start the Next.js application
CMD ["pnpm", "run", "dev"]
