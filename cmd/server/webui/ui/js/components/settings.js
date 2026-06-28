import { api } from '../api.js';
import { state } from '../state.js';
import { notifications } from '../utils/notifications.js';
import { t } from '../utils/i18n.js';

const defaultFailover = {
    recoveredEndpointPolicy: 'deprioritize',
    cooldowns: {
        quotaExhaustedSec: 3600,
        rateLimitedSec: 120,
        upstreamErrorSec: 60,
        networkErrorSec: 30,
        tokenUnavailableSec: 600,
        configErrorSec: 1800
    },
    circuitBreaker: {
        consecutiveFailures: 3,
        windowSec: 60,
        failureRateThreshold: 0.60,
        minRequests: 5,
        cooldownSec: 600
    }
};

class Settings {
    constructor() {
        this.container = document.getElementById('view-container');
        this.currentFailover = defaultFailover;
        window.addEventListener('languageChanged', () => {
            if (state.get('currentView') === 'settings') {
                this.render();
            }
        });
    }

    async render() {
        this.container.innerHTML = `
            <div class="settings">
                <h1>${t('settings.title')}</h1>
                <div class="card mt-3">
                    <div class="card-header">
                        <h3 class="card-title">${t('settings.failoverTitle')}</h3>
                    </div>
                    <div class="card-body">
                        <form id="settings-form">
                            <div class="form-group">
                                <label class="form-label">${t('settings.recoveredEndpointPolicy')}</label>
                                <select class="form-select" name="recoveredEndpointPolicy">
                                    <option value="deprioritize">${t('settings.policies.deprioritize')}</option>
                                    <option value="auto_return">${t('settings.policies.autoReturn')}</option>
                                </select>
                            </div>
                            <div class="grid grid-cols-2">
                                ${this.renderCooldownInput('quotaExhaustedSec', t('settings.cooldowns.quotaExhausted'))}
                                ${this.renderCooldownInput('rateLimitedSec', t('settings.cooldowns.rateLimited'))}
                                ${this.renderCooldownInput('upstreamErrorSec', t('settings.cooldowns.upstreamError'))}
                                ${this.renderCooldownInput('networkErrorSec', t('settings.cooldowns.networkError'))}
                                ${this.renderCooldownInput('tokenUnavailableSec', t('settings.cooldowns.tokenUnavailable'))}
                                ${this.renderCooldownInput('configErrorSec', t('settings.cooldowns.configError'))}
                            </div>
                            <button type="submit" class="btn btn-primary mt-3">${t('common.save')}</button>
                        </form>
                    </div>
                </div>
            </div>
        `;

        document.getElementById('settings-form').addEventListener('submit', (event) => this.save(event));
        await this.load();
    }

    renderCooldownInput(name, label) {
        return `
            <div class="form-group">
                <label class="form-label">${label}</label>
                <input class="form-input" type="number" min="0" name="${name}">
            </div>
        `;
    }

    async load() {
        try {
            const config = await api.getConfig();
            const failover = this.normalizeFailover(config.failover);
            this.currentFailover = failover;
            const form = document.getElementById('settings-form');
            form.elements.recoveredEndpointPolicy.value = failover.recoveredEndpointPolicy;
            Object.entries(failover.cooldowns).forEach(([key, value]) => {
                if (form.elements[key]) {
                    form.elements[key].value = value;
                }
            });
        } catch (error) {
            notifications.error(`${t('settings.failedToLoad')}: ${error.message}`);
        }
    }

    async save(event) {
        event.preventDefault();
        const form = event.currentTarget;
        const readSeconds = (name) => {
            const value = Number.parseInt(form.elements[name]?.value || '0', 10);
            return Number.isFinite(value) && value > 0 ? value : 0;
        };

        try {
            await api.updateConfig({
                failover: {
                    recoveredEndpointPolicy: form.elements.recoveredEndpointPolicy.value,
                    cooldowns: {
                        quotaExhaustedSec: readSeconds('quotaExhaustedSec'),
                        rateLimitedSec: readSeconds('rateLimitedSec'),
                        upstreamErrorSec: readSeconds('upstreamErrorSec'),
                        networkErrorSec: readSeconds('networkErrorSec'),
                        tokenUnavailableSec: readSeconds('tokenUnavailableSec'),
                        configErrorSec: readSeconds('configErrorSec')
                    },
                    circuitBreaker: this.currentFailover?.circuitBreaker || defaultFailover.circuitBreaker
                }
            });
            notifications.success(t('settings.saved'));
            await this.load();
        } catch (error) {
            notifications.error(`${t('settings.failedToSave')}: ${error.message}`);
        }
    }

    normalizeFailover(failover) {
        const cooldowns = failover?.cooldowns || {};
        const circuitBreaker = failover?.circuitBreaker || {};
        return {
            recoveredEndpointPolicy: failover?.recoveredEndpointPolicy || defaultFailover.recoveredEndpointPolicy,
            cooldowns: {
                quotaExhaustedSec: Number.isFinite(Number(cooldowns.quotaExhaustedSec)) ? Number(cooldowns.quotaExhaustedSec) : defaultFailover.cooldowns.quotaExhaustedSec,
                rateLimitedSec: Number.isFinite(Number(cooldowns.rateLimitedSec)) ? Number(cooldowns.rateLimitedSec) : defaultFailover.cooldowns.rateLimitedSec,
                upstreamErrorSec: Number.isFinite(Number(cooldowns.upstreamErrorSec)) ? Number(cooldowns.upstreamErrorSec) : defaultFailover.cooldowns.upstreamErrorSec,
                networkErrorSec: Number.isFinite(Number(cooldowns.networkErrorSec)) ? Number(cooldowns.networkErrorSec) : defaultFailover.cooldowns.networkErrorSec,
                tokenUnavailableSec: Number.isFinite(Number(cooldowns.tokenUnavailableSec)) ? Number(cooldowns.tokenUnavailableSec) : defaultFailover.cooldowns.tokenUnavailableSec,
                configErrorSec: Number.isFinite(Number(cooldowns.configErrorSec)) ? Number(cooldowns.configErrorSec) : defaultFailover.cooldowns.configErrorSec
            },
            circuitBreaker: {
                consecutiveFailures: Number.isFinite(Number(circuitBreaker.consecutiveFailures)) ? Number(circuitBreaker.consecutiveFailures) : defaultFailover.circuitBreaker.consecutiveFailures,
                windowSec: Number.isFinite(Number(circuitBreaker.windowSec)) ? Number(circuitBreaker.windowSec) : defaultFailover.circuitBreaker.windowSec,
                failureRateThreshold: Number.isFinite(Number(circuitBreaker.failureRateThreshold)) ? Number(circuitBreaker.failureRateThreshold) : defaultFailover.circuitBreaker.failureRateThreshold,
                minRequests: Number.isFinite(Number(circuitBreaker.minRequests)) ? Number(circuitBreaker.minRequests) : defaultFailover.circuitBreaker.minRequests,
                cooldownSec: Number.isFinite(Number(circuitBreaker.cooldownSec)) ? Number(circuitBreaker.cooldownSec) : defaultFailover.circuitBreaker.cooldownSec
            }
        };
    }
}

export const settings = new Settings();
