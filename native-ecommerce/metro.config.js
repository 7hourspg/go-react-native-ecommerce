// This is the default metro.config.js file

// const { getDefaultConfig } = require('expo/metro-config');
// const { withNativeWind } = require('nativewind/metro');

// const config = getDefaultConfig(__dirname);

// module.exports = withNativeWind(config, { input: './global.css', inlineRem: 16 });

// This is the modified metro.config.js file to exclude Stripe on web
const { getDefaultConfig } = require('expo/metro-config');
const { withNativeWind } = require('nativewind/metro');

const config = getDefaultConfig(__dirname);

// Add resolver to exclude Stripe on web
config.resolver.resolveRequest = (context, moduleName, platform) => {
  if (platform === 'web' && moduleName === '@stripe/stripe-react-native') {
    // Return empty module for web
    return {
      type: 'empty',
    };
  }
  // Use default resolution for everything else
  return context.resolveRequest(context, moduleName, platform);
};

module.exports = withNativeWind(config, { input: './global.css', inlineRem: 16 });