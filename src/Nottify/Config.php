<?php
/**
 * Provide Configuration functionality.
 */
namespace Nottify;

/**
 * Simple Config detail for ease of customisation.
 *
 * @package Nottify
 * @author Allan Shone <allan.shone@yahoo.com>
 */
class Config
{
    /**
     * Load base Configuration file.
     *
     * @param string $filename
     *  Filename for the Configuration content.
     */
    public function __construct($filename)
    {
        $this->filename = $filename;

        $this->digest();
    }

    /**
     * Handle dynamic calls.
     *
     * @param string $name
     *  Name of the method being called.
     * @param array $arguments
     *  All provided arguments.
     * @return mixed
     * @throws \InvalidArgumentException
     */
    public function __call($name, $arguments)
    {
        $config_name = array_shift($arguments);
        if (!isset($this->config[$config_name])) {
            throw new \InvalidArgumentException('Config entry not found');
        }

        $type = substr($name, 3);
        if (substr($name, 0, 3) === 'get') {
            switch (strtolower($type)) {
                case 'bool':
                    return (bool)($this->config[$config_name]);
                case 'string':
                    return strval($this->config[$config_name]);
                case 'int':
                case 'integer':
                    return intval($this->config[$config_name]);
                case 'csv':
                    return explode(',', $this->config[$config_name]);
            }

            throw new \InvalidArgumentException('Invalid type selection');
        } else if (substr($name, 0, 2) === 'is') {
            switch (strtolower($type)) {
                case 'bool':
                    return is_bool($this->config[$config_name]);
                case 'string':
                    return is_string($this->config[$config_name]);
                case 'int':
                case 'integer':
                    return is_int($this->config[$config_name]);
            }

            throw new \InvalidArgumentException('Invalid type selection');
        }
    }

    /**
     * Determine the Slim-required Mode of operation.
     *
     * @param bool $debug
     *  If debug mode is enabled.
     * @return string
     */
    public function determineMode($debug)
    {
        if ($debug === true) {
            return 'development';
        }

        return 'production';
    }

    /**
     * Determine if the browser is valid.
     *
     * @param array $input
     *  Possible content for user determining.
     * @return bool
     */
    public function isValidUser($user)
    {
        if (isset($user['HTTP_CLIENT_IP']) ||
            isset($user['HTTP_X_FORWARDED_FOR']) ||
            !isset($user['REMOTE_ADDR'])) {
            return false;
        }

        $valid_ip = $this->getCsv('main.ips');
        if (!in_array($user['REMOTE_ADDR'], $valid_ip)) {
            return false;
        }

        return true;
    }

    /**
     * Parse and import the configuration.
     *
     * @throws \InvalidArgumentException
     * @throws \RuntimeException
     */
    private function digest()
    {
        if (!file_exists($this->filename)) {
            throw new \InvalidArgumentException('Configuration file does not exist');
        }

        $config = parse_ini_file($this->filename, true);

        if (empty($config)) {
            throw new \RuntimeException('Configuration is empty');
        }

        foreach ($config as $section => $specific) {
            foreach ($specific as $name => $value) {
                $this->config["{$section}.{$name}"] = $value;
            }
        }
    }

    /**
     * Store the actual configuration content.
     *
     * @var array
     */
    private $config = array();

    /**
     * Configuration content file.
     *
     * @var string
     */
    private $filename = '';
}

