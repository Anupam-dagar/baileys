# Configuration Management

Baileys offers a robust configuration management system powered by [viper](https://github.com/spf13/viper), simplifying the process of adapting your server to different environments. By adhering to a few conventions and utilizing predefined keys, you can effortlessly tailor your server's behavior. With convenient helper functions at your disposal, retrieving configuration values becomes a seamless task, ensuring your application operates smoothly and efficiently under varying conditions. Baileys empowers developers with the flexibility needed for effective configuration management.

## Setting up configuration
It is essential to establish a `config` directory at the project's root, housing a configuration file named `dev.yaml`. The filename corresponds to the value of the `ENV` environment variable.

## Pre defined configuration keys
Baileys employs predefined keys, enabling effortless server setup for various operations. For instance, it streamlines database connectivity. Below is a list of these keys and their intended purposes:
- `database`: Contains configuration settings for connecting to the database.
  - `host`: Specifies the database's host URL.
  - `port`: Port at which the database is running.
  - `username`: Specifies the username used for database connectivity.
  - `password`: Specifies the password user for database connectivity.
  - `name`: Name of the database to connect to.
- `server`: Server level configuration parameters
  - `port`: Specifies the port on which the server will run.
  - `base_api_path`: Defines the base path for all API endpoints.

## Helper functions which can be used for development
The configuration package offers a set of handy methods for retrieving configuration values. Here are more details on these functions:
1. `GetConfiguration` - Returns the pointer to `viper` configuration object.
2. `GetStringConfig` - Returns the string value for the provided key.
3. `GetIntConfig` - Returns the int value for the key provided key.
4. `GetBoolConfig` - Returns the bool value for the key provided key.
5. `GetConfig` - Returns an interface{} value for the supplied key, which can be type-asserted to the user-defined type as needed.

